#!/bin/bash

go build -o InrixBackend cmd/*.go && ./InrixBackend -dbname=InrixBackend -dbuser=rahuldhiman