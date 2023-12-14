# The application is crashing

## Details

This normally occurs when there is a resource leak of some description.

## Impact

The application crashing will mean downtime for our customers.


## Mitigation

0. ssh into the server

   ```text
   ssh hostname@ip
   ```

1. restart the application

   ```text
   sudo systemctl restart runbooks
   ```
