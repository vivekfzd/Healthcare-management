## download the minifab images in linux
mkdir -p ~/mywork && cd ~/mywork && curl -o minifab -sL https://tinyurl.com/yxa2q6yr && chmod +x minifab

## go to mywork folder

## database is couchdb and expose the endpoint
./minifab netup -s couchdb -e true

## create  a channel
./minifab create -c channel1

## join th peer
./minifab join

## see the configration of channel
./minifab channelquery

## provide the access to owner
sudo chown -R <owner_name> /home/ec2-user/mywork

change the batch timeout is 0.25 sec in channel1_configration.json file

## channel signoff and update
./minifab channelsign,channelupdate

## copy the healthcare folder and paste to vars/chaincode folder

## install the chaincode
./minifab install -n healthcare -l go

## approve, commit, anchorupdate and discover command
./minifab approve,commit,anchorupdate,discover

## init the chaincode
./minifab initialize -p '"Init"'

## if you modify the chaincode then update the chaincode 
./minifab ccup -n healthcare -l go -v 2.0

-----------------------------------------------------------------------------------------
## create patien command
create Patient
./minifab invoke -p '"Admin","createPatient","P1","vivek","23","M","O","balla","000000000"'

## get the patient
./minifab invoke -p '"Admin","getPatient","P1"'

'{\\"id\\":\\"P1\\",\\"name\\":\\"xx\\",\\"age\\":23,\\"gender\\":\\"M\\",\\"bloodType\\":\\"O\\",\\"address\\":\\"Balla\\",\\"phoneNumber\\":\\"0000000000\\",\\"recordIDs\\":[]}'

----------------------------------------------------------------------------------------

## delete the patient 
./minifab invoke -p '"Admin","deletePatient","P1"'

--------------------------------------------------------------------------------------------------

## Create Doctor
./minifab invoke -p '"Admin","createDoctor","D10","sp rai","ENT","0123456789"'

## query the doctor
./minifab invoke -p '"Admin","getDoctor","D10"'

'{\\"id\\":\\"D1\\",\\"name\\":\\"sp rai\\",\\"specialty\\":\\"gyno\\",\\"phoneNumber\\":\\"0123456789\\"}'

--------------------------------------------------------------------------------------------------

## grant write access to doctor
./minifab invoke -p '"Patient","grantWriteAccess","P10","D10"'


## read patient
./minifab invoke -p '"Admin","getPatient","P1"' -> Admi./minifab initialize -p '"Init"'n to patient


------------------------------------------------------------------------------------------------------

## Create Medical Record

./minifab invoke -p '"Doctor","createMedicalRecord","M10","P10","D10","05-05-2023","Percrespition1"'


./minifab invoke -p '"Doctor","createMedicalRecord","M2","P1","D1","13-05-2023","Percrespition2"'


------------------------------------------------------------------------------------
## check the medical record is added or not

./minifab invoke -p '"Patient","getPatient","P1"'
'{\\"id\\":\\"P1\\",\\"name\\":\\"xx\\",\\"age\\":23,\\"gender\\":\\"M\\",\\"bloodType\\":\\"O\\",\\"address\\":\\"Balla\\",\\"phoneNumber\\":\\"0000000000\\",\\"recordIDs\\":[\\"1\\"]}' 


----------------------------------------------------------------------------------------------------
## read all the medical records
./minifab invoke -p '"Patient","readMedicalRecords","P10"'

'[{\\"id\\":\\"M1\\",\\"patientID\\":\\"P1\\",\\"doctorID\\":\\"D1\\",\\"date\\":\\"05-05-2023\\",\\"prescription\\":\\"Percrespition1\\"},{\\"id\\":\\"M2\\",\\"patientID\\":\\"P1\\",\\"doctorID\\":\\"D1\\",\\"date\\":\\"13-05-2023\\",\\"prescription\\":\\"Percrespition2\\"}]'

----------------------------------------------------------------------------------------------

## create Labtechnician
./minifab invoke -p '"Admin","createLabTechnician","T1","vivek"'

./minifab invoke -p '"Admin","getLabTechnician","T1"'


------------------------------------------------------------------------------------------------------
./minifab invoke -p '"Admin","createInsuranceCompany","C1","vinay","xyz"'

./minifab invoke -p '"Admin","getInsuranceCompany","C1"'

-----------------------------------------------------------------------------------------------------------------------
