## Create a Gandalf Connector

Design Steps:

0. Declare Major/Minor (int64)
1. Pick up Input/Outputs
2. Declare new Worker instance part
3. OAuth and context configure
4. "RegisterCommandFunc()" to create according to commands you want
5. Create functions acccording to commands

### Other parts

**Major/Minor:** Update "Major.Minor" version
    
    v1.0

    Version Major.Minor

**Payload (from function):** Arguments from command

- - - 

## Described Stages

### Pick Up Input/Output

Get scan/input and display it, in Golang :

    input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println(input.Text())
