#!/bin/bash
set -euo pipefail

if [[ ! -f "./go.mod" ]]; then
    echo "Need to run this from the directory containing go.mod"
    exit 1
fi

OUTPUT_FILE="ACKNOWLEDGEMENTS.md"

# Ensure modules are downloaded
go mod download

# Get go-licenses tool
go install github.com/google/go-licenses@latest
GO_LICENSES="${HOME}/go/bin/go-licenses"

# Save licenses
export TEMPDIR="$(mktemp -d /tmp/generate-acknowledgements.XXXXXX)" || exit 1
rmdir "${TEMPDIR}"
"${GO_LICENSES}" save . --save_path="${TEMPDIR}"
trap "rm -fr ${TEMPDIR}" EXIT

# Build acknowledgements file
export TEMPFILE="$(mktemp acknowledgements.XXXXXX)" || exit 1
cat > "${TEMPFILE}" <<EOF
# Buildkite Agent OSS Attributions

The Buildkite Agent would not be possible without open-source software. 
Licenses for the libraries used are reproduced below.
EOF

addfile() {
    printf "\n\n---\n\n## %s\n\n\`\`\`\n" "${2:-${1#${TEMPDIR}/}}" >> "${TEMPFILE}"
    cat "$1" >> "${TEMPFILE}"
    echo "\`\`\`" >> "${TEMPFILE}"
}

## The Go standard library also counts.
addfile "$(go env GOROOT)/LICENSE" "Go standard library"

## Now add all the modules.
export -f addfile
find "${TEMPDIR}" -type f -exec bash -c 'addfile "$1"' _ {} \;

## Add trailer
printf "\n\n---\n\nFile generated using %s\n%s\n" "$0" "$(date)" >> "${TEMPFILE}" 

# Replace existing acknowledgements file
mv "${TEMPFILE}" "${OUTPUT_FILE}"
chmod 644 "${OUTPUT_FILE}"

# gzipped version for embedding purposes
gzip -kf "${OUTPUT_FILE}"

exit 0
