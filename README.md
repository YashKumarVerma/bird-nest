# Bird Nest :bird: :tiger:

![Build package](https://github.com/YashKumarVerma/bird-nest/workflows/Build%20package/badge.svg?branch=repl)


Bird Nest is a cli to quickly generate Nest.Js modules by defining the schema in a sql inspired markup language (don't worry, we've got autocomplete to assist you). The utility generates data transfer objects, entity, repository, service, module and controller files based on the schema provided.

## CLI
![https://i.imgur.com/7EA83W9.png](https://i.imgur.com/7EA83W9.png)
![https://i.imgur.com/iUK7Ytn.png](https://i.imgur.com/iUK7Ytn.png)
![https://i.imgur.com/qKAVnxx.png](https://i.imgur.com/qKAVnxx.png)

When done, enter **exit**
![https://i.imgur.com/Nk3W1Sg.png](https://i.imgur.com/Nk3W1Sg.png)

## Output
The generated files have this structure.

![https://i.imgur.com/kTDdKUt.png](https://i.imgur.com/kTDdKUt.png)

To view the files generated, please check the **`generated`** directory.

## Usage
you need to have have golang setup on your system, as till the time we hit v1.0, we won't be releasing any binaries. To run the current version of the script, do the following:

```bash
# clone the repository
git clone https://github.com/YashKumarVerma/bird-nest

# change directory
cd bird-nest

# now you can build a binary for yourself, or directly run
# Building
go build .

# or running directly
go run main.go
```
You might want to clear the contents of **`generated`** directory. Note that files with same module name only are overwritten.

## Development
It's the v0.0 of the project, will mode towards v1.0 soon. Help us by opening an issue, or get your hands dirty and send a pull request instead.