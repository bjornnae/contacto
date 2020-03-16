# Contacto

Contact tester.

## Develop

For local development and testing, a test server is bundled. Start this so that the test code in contacto_test.go has anything to call.

## Run code

    $ go run .

## Build code into executable

    $ go build .
    
## Build code into executable for other platform than the development OS

    $ GOOS=windows go build .


## Running executable standalone

1. Put the file testCatalogue.txt in the same folder as the executable contacto.exe .

2. Update the lines in the file testCatalogue.txt so that every line follow the form:

    'hostname url'

    Delimiter between hostname and url must be SPACE character.

3. run the file 'contacto.exe' 
