#!/bin/bash

# Variables
ETCD_CERT=/etc/etcd/ssl/etcd.pem
ETCD_KEYS=/etc/etcd/ssl/etcd-key.pem
ETCD_CACERT=/etc/kubernetes/ssl/ca.pem

ETCD_DATA_DIR=/var/lib/etcd #only required for etcdv3
ETCD_BACKUP_DIRECTORY=/data/etcd_backup
ETCD_BACKUP_DATATIMES=$(date "+%Y%m%d_%H%M%S")
ENDPOINTS=https://192.168.133.128:2379,https://192.168.133.129:2379,https://192.168.133.130:2379


# create the backup directory if it doesn't exist
[[ -d $ETCD_BACKUP_DIRECTORY ]] || mkdir -p $ETCD_BACKUP_DIRECTORY
    
# backup etcd v3 data
  export ETCDCTL_API=3
  /usr/bin/etcdctl \
      --cacert=$ETCD_CACERT \
      --cert=$ETCD_CERT \
      --key=$ETCD_KEYS \
      --endpoints=$ENDPOINTS \
      snapshot save $ETCD_BACKUP_DIRECTORY/${ETCD_BACKUP_DATATIMES}_snapshot.db


# check if backup interval is set
if [[ ! -f "$ETCD_BACKUP_DIRECTORY/${ETCD_BACKUP_DATATIMES}_snapshot.db" ]]; then
    echo "$(date +"%Y-%m-%d %H:%M:%S")	  Snapshot backup failed!" >> /var/log/backup_etcd.log
    exit 1
else
    echo "$(date +"%Y-%m-%d %H:%M:%S")	  The snapshot backup was successful. Please go to the $ETCD_BACKUP_DIRECTORY directory to check the backup file!"  >> /var/log/backup_etcd.log
fi
