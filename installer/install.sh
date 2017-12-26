#!/bin/bash

sudo tar -xzvf rootfiles.tar.gz -C /
sudo systemctl start rpilight.service
sudo systemctl enable rpilight.service
