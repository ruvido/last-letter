# error handling
- panics if template folder is not there

# dev environment to develop letter
- create a dev branch
- make code modifications
- test new executable
- rename current master to master-RELEASE (eg. 0.1)
- rename dev to master
prompt: what is the best strategy to have a dev branch in git, gimme step to follow., eg: i am in the master branch but now i want to start modifications, what to do? create a dev branch? Once i have finished to do modifications, what the best way to update master? rename dev to master? make a merge? gimme the most robust/easy strategy to develop in git.

# Automatic sendind emails based on date
- before sending check if the date is today
- add a flag --now to send the newsletter now
- test/debug: send the newsletter in anycase (dont check the date)
- make a cron job to start letter everyday

# cron setup for letter
1. just make a cron job
	pro: very easy, predictable
	con: not robust in case of migration, bound to a specific user
2. make a docker file with a cronjob
	pro: easy migration, clean
	con: more difficult setup

# prompt to create a docker file
write a docker file to launch an executable called letter. abort with an error if the a config file called "config.toml" is not present in the present folder. check if the executable "letter" is in the current folder, otherwise download it from github at the address github.com/ruvido/letter and compile it, it's golang. put the source code in a subfolder "src" of the present directory. now once the config and the executable are present in the current folder, i'd like that the executable is launched according to a cron rule (eg everyday at 5am)

# admin email results
An email should be sent to admin with the results of the newsetter send:
	- number of addresses
	- number of bounces

# change ruvido email to -> admin in pocketbase
