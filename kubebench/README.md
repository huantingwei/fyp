## Usage

1. run `gcloud auth revoke` to remove existing authentication data

2. run `./kubebench.sh`

3. Enter the cluter name, zone name (e.g. us-central1-a) and project name

4. Use the URL in stdout or url.txt to get the verification code. Open a new text file in the current directory, paste to code, save it as token.txt

5. After awhile, the kubebench.txt will be saved in the current directory, containing the auditing results for section 4 to 6