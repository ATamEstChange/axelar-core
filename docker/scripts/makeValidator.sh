#!/bin/bash

scavengeCLI tx scavenge transferTokens "$(scavengeCLI keys show validator -a)" 100000000stake \
  --from treasury --yes

sleep 3

scavengeCLI tx staking create-validator --yes\
 --amount 100000000stake \
 --moniker "${1}" \
 --commission-rate="0.10" \
 --commission-max-rate="0.20" \
 --commission-max-change-rate="0.01" \
 --min-self-delegation="1" \
 --pubkey "$(scavengeD tendermint show-validator)" \
 --from validator
