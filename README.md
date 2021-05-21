# Webdav File Revisioner
Makes new revisions of specified file on webdav drive.

## Usage
### Config
Copy config.json.sample to config.json and fill out the form.

|Field Name|Description|Example|
|----------|-----------|-------|
|FilePath|Path to file to make revisions of.|folder/file.txt|
|DestinationPath|Path to folder to put backups. Folder will be saved with original file name and extension with the date and time appended. Make sure to include double slashes or this may just put the file at the root with the folder names prepended to the file. Ensure that there are slashes at the end of the path too.|folder1//folder2//|
|WebdavUrl|The url of the webdav server.|http://example.com/|
|WebdavUser|The username to log into the webdav server.|example@example.com|
|WebdavPassword|The password to log into the webdav server.|Password123|
|UseEmailer|Whether or not the program should email when an error occurs during backup.|true|
|FromEmail|The email address the error emails get sent from|example@example.com|
|FromEmailPassword|The from email address's password.|Password123|
|ToEmail|The email address to send the error emails to.|example@example.com|
|SmtpSever|The smtp server for the email service you're using.|smtp.example.com:587|
|SmtpAuthServer|The auth server for the email service you're using.|smtp.example.com|

### Build Docker Image
Run the following command to build the docker image.
```shell
docker build -t webdav-file-revisioner .
```

### Run Docker Container
The following can be used to run the docker container. The rm flag is used to delete the container immediately upon exiting the application.
```shell
docker run -d --rm webdav-file-revisioner
```
This will only run the copy process once. To further automate this a cron job can be configured to run once per day at midnight with the following in your crontab file.
```shell
0 0 * * * docker run -d --rm webdav-file-revisioner
```