# get into project folder

graphql-apollo-api

# create a .env file inside root project
# put the following examples inside with valid values

PORT=3001
MONGO_DB_URL='mongodb+srv://user:password@clustername.cvepz.mongodb.net/<dbname>?retryWrites=true&w=majority'
JWT_SECRET_KEY=test@1234

# install dependencies

 npm install

# Run local dev server

npm run dev

# When using the graphql playground -

After create the user, you should login to create the tasks
put your email and password to continue.
You will receive a token, with this token you can : Create Tasks, Update Tasks, Delete Tasks and Change Task Status

Put the token inside http Headers :

Example :

{
	"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZsb3JhQGZsb3JhLmNvbSIsImlhdCI6MTU5Nzg3NjU3MCwiZXhwIjoxNTk4MTM1NzcwfQ.0ny0ySjezeITjDfFZcIOPJIKTHqB2HhxgESeqZZ8xEs"
}
