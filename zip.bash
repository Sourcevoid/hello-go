#!/bin/bash

# Remove old archive
rm -rf archive.zip

# Create archive
zip -r -q archive.zip . -x .git/\*

