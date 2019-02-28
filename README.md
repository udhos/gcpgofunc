# gcpgofunc

# Install GCP SDK

See: https://cloud.google.com/sdk/docs/quickstart-linux

Recipe:

    wget https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-231.0.0-linux-x86_64.tar.gz
    tar xf google-cloud-sdk-231.0.0-linux-x86_64.tar.gz
    ./google-cloud-sdk/install.sh
    gcloud components update
    gcloud components install beta
    gcloud init

# Grab gcpgofunc

    git clone https://github.com/udhos/gcpgofunc

# Deploy with an HTTP trigger

    cd gcpgofunc/helloworld
    gcloud functions deploy HelloGet --runtime go111 --trigger-http

# Get information about deployed function

    gcloud functions describe HelloGet

# Call the HTTP trigger

    curl -d '{"name":"bogus"}' https://xxxxx.cloudfunctions.net/HelloGet

# GCP Functions Quickstart

Quickstart: https://cloud.google.com/functions/docs/quickstart
