# The application is crashing

## Details

This normally occurs when there is a resource leak of some description.

## Impact

The application crashing will mean downtime for our customers.


## Mitigation

### When

To execute the mitigation steps, the condition below must evaluate to `true` first.

```
echo 1;
```

### Steps

Below are the steps that must be performed in order to mitigate this incident.

0. ssh into the server

   ```text
   ssh hostname@ip
   ```

1. restart the application

   ```text
   sudo systemctl restart runbooks
   ```

   Verify the command completed as expected, asserting the below command returns `true` :-

    ```text
    if [ -d "/tmp/runbooks-cli" ];
    then 
      echo "false"
    else 
      echo "true"
    fi
    ```
