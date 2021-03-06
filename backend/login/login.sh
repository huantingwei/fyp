#!/bin/bash
#-------Part 0: revoke all gcloud accounts-------------
if [ $FYPENV = "prod" ]; then
	gcloud auth revoke --all
	echo hello
fi

#-------Part 1: automated authentication, connect gcloud to user's cloud project--------
#-------as well as connect kubectl to the GKE cluster in that project--------

#all stdout will be written to output.txt too
stdout="./output.txt"

#The OAuth URL for Google Cloud authentication will be written into url.txt
url="./url.txt"

#token.txt will store the access token generated from the OAuth authentication process
token="./token.txt"

#for string matching in output.txt, to detect the appearance of the OAuth URL
string="oauth2"

#projindex.txt will contain the index number which matches the project name above
projindex="./projindex.txt"

# ask for project name, cluster name and GCP project name
read -p 'GKE cluster name: ' clustername
read -p 'GCP project name: ' projectname
read -p 'GKE zone name: ' zonename

function loop {
	#reinitialise gcloud [default] configuration
	echo "1"
	#enter 'Y' to login
	echo "Y"
	
	#loop until the OAuth URL appears
	while true
	do
		#True if the OAuth URL appears in stdout
		if [ -e $stdout ] && [ ! -z $(grep "$string" "$stdout") ]; then 
		
			#extract the URL from stdout, write it to url.txt
			grep $string $stdout | xargs > $url
			break
		else
			sleep 5
		fi
	done
	
	#loop until the access token is passed by from user and written into token.txt		
	while true
	do
		# true if token.txt exists
		if [ -s $token ]; then
			while read line; do
			echo $line
			done < $token
			
			break
		else
			sleep 5
		fi
	done
	
	while true
	do
		# true if the project options appears in stdout
		if [ -e $stdout ] && [[ ! -z $(grep "$projectname" "$stdout") ]]; then
			
			# extract the desired project option from stdout, write it to projindex.txt
			grep $projectname $stdout | xargs > $projindex
			
			# extract the index number which is the third char
			while read line; do
			# echo the third char of line
			echo ${line:1:1}
			done < $projindex
			
			echo 'n'
			
			break
	
		else
			sleep 5
		fi
	
	done
	
}

# 2>&1 | tee will write a copy of stdout to output.txt
loop | gcloud init --console-only --skip-diagnostics 2>&1 | tee ./output.txt

#connect kubectl to the target GKE cluster
gcloud container clusters get-credentials $clustername --zone $zonename --project $projectname

sleep 10

echo "Finish authentication"

#cleanup
rm $stdout $url $token $projindex