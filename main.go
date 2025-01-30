package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	mf "github.com/Asri-Mohamad/Master_Function"
)

type dataUser struct {
	Name     string `json:"name"`
	Family   string `json:"family"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

var fileName string = "Userdata.json"

func main() {
	for {
		com := showMenu()
		switch com {
		case 1:
			fmt.Println("....<<<< Register Form >>>>....")
			registerForm()
		case 2:
			fmt.Println("....<<<< Login Form >>>>....")
			loginForm()

		case 3:
			fmt.Println("....<<<< Edivart Form >>>>....")
			editForm()

		case 4:
			fmt.Println("....<<<< Delete Form >>>>....")
			deleteForm()

		case 5:
			return
		}

	}
}

// ---------------------------- load file ------------------------------
func loadFile(fileName string) ([]dataUser, int) {
	newReadUserPass := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your User name :")
	userName, _ := newReadUserPass.ReadString('\n')
	userName = strings.TrimSpace(userName)
	fmt.Print("Enter your Password :")
	pass, _ := newReadUserPass.ReadString('\n')
	pass = strings.TrimSpace(pass)

	var readStruct []dataUser
	var index int = -1
	var fileOpen bool = true
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0)
	if err != nil {
		fmt.Println("Open data file have problem ...... ")

		fileOpen = false
	} else {
		defer file.Close()
		allData, err := io.ReadAll(file)
		if err != nil {
			fmt.Println("Read data from file have problem....")

			fileOpen = false
		} else {
			err = json.Unmarshal(allData, &readStruct)
			if err != nil {
				fmt.Println("Convert to jason file have problem....")

				fileOpen = false
			}

		}

	}
	if fileOpen {
		mach := false
		for i, user := range readStruct {

			if user.UserName == userName {
				if user.Password == pass {
					index = i
					mach = true
					break
				}
			}

		}
		if !mach {
			index = -2 // login not mach
		}
	} else {
		index = -1 // open file problem
	}
	return readStruct, index
}

// ----------------------------Delete form -----------------------------
func deleteForm() {

	readStruct, index := loadFile(fileName)
	if index == -2 {
		fmt.Println("The user name Or password is worng...")
	} else {
		fmt.Printf("\nYour information is ...:\n   Name : %s\n   Family :%s\n   User name : %s\n Do you want to delete (y/n)?",
			readStruct[index].Name, readStruct[index].Family, readStruct[index].UserName)
		if yesOrNo() {

			readStruct = append(readStruct[:index], readStruct[index+1:]...)
			file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
			if err != nil {
				fmt.Println("Open data file have problem...")

			} else {
				file.Truncate(0)
				file.Seek(0, 0)
				encode := json.NewEncoder(file)
				err = encode.Encode(readStruct)
				if err != nil {
					fmt.Println("problem to save data ....")

				} else {
					fmt.Printf("Yes\n ID Deleted ....")
				}

			}

		} else {
			fmt.Println("No\nThe ID not deleted ...")
		}
	}
	_ = mf.CharGetKey()
}

// ---------------------------- Edit form -----------------------------
func editForm() {

	readStruct, index := loadFile(fileName)

	if index == -2 {
		fmt.Println("The user name Or password is worng...")
	} else {
		var newUser dataUser
		fmt.Printf("\nYour information is ...:\n   Name : %s\n   Family :%s\n   User name : %s\n Pleas Enter new data ....\n",
			readStruct[index].Name, readStruct[index].Family, readStruct[index].UserName)

		newReadinformation := bufio.NewReader(os.Stdin)
		fmt.Print("New Name :")
		newUser.Name, _ = newReadinformation.ReadString('\n')
		newUser.Name = strings.TrimSpace(newUser.Name)
		fmt.Print("New Family :")
		newUser.Family, _ = newReadinformation.ReadString('\n')
		newUser.Family = strings.TrimSpace(newUser.Family)
		fmt.Print("New User name : ")
		newUser.UserName, _ = newReadinformation.ReadString('\n')
		newUser.UserName = strings.TrimSpace(newUser.UserName)
		for {
			fmt.Print("Enter new password:")
			newUser.Password, _ = newReadinformation.ReadString('\n')
			newUser.Password = strings.TrimSpace(newUser.Password)
			fmt.Print("Enter new password again:")
			repass, _ := newReadinformation.ReadString('\n')
			repass = strings.TrimSpace(repass)
			if newUser.Password != repass {
				println("Password and Repassword is not mach plese Enter again...")
			} else {
				newUser.Password = repass
				break
			}
		}

		fmt.Print("\n Do you want save new data (y/n)?")
		if yesOrNo() {
			readStruct[index] = newUser
			file, err := os.OpenFile("Userdata.json", os.O_RDWR, 0666)
			if err != nil {
				fmt.Println("Open data file have problem...")

			} else {
				defer file.Close()
				file.Truncate(0)
				file.Seek(0, 0)
				encode := json.NewEncoder(file)
				err = encode.Encode(readStruct)
				if err != nil {
					fmt.Println("problem to save data ....")

				} else {
					fmt.Printf("Yes\n Data Saveing ....")
				}

			}

		} else {
			fmt.Println("No\n Last informatin not changing....")
		}
	}

	_ = mf.CharGetKey()

}

// --------------------------- Login form -----------------------------

func loginForm() {
	readStruct, index := loadFile(fileName)
	if index == -2 {

		fmt.Println("The user name Or password is worng...")

	} else {
		fmt.Printf("Welcome %s %s to the system....", readStruct[index].Name, readStruct[index].Family)

	}
	_ = mf.CharGetKey()

}

// ------------------------- Register form ----------------------------
func registerForm() {
	var newUser dataUser
	newReadData := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Name:")
	name, _ := newReadData.ReadString('\n')
	newUser.Name = strings.TrimSpace(name)

	fmt.Print("Enter Family:")
	family, _ := newReadData.ReadString('\n')
	newUser.Family = strings.TrimSpace(family)

	fmt.Print("Enter User Name:")
	userName, _ := newReadData.ReadString('\n')
	newUser.UserName = strings.TrimSpace(userName)

	for {
		fmt.Print("Enter password:")
		pass, _ := newReadData.ReadString('\n')
		pass = strings.TrimSpace(pass)

		fmt.Print("Enter password again:")
		repass, _ := newReadData.ReadString('\n')
		repass = strings.TrimSpace(repass)

		if pass != repass {
			println("Password and Repassword is not mach plese Enter again...")
		} else {
			newUser.Password = repass
			break
		}
	}

	fmt.Printf("This user was create:\n	Name:%s\n	Family:%s\n	Username:%s\n	Password:%s\n Are you shor to save this data(Y/N)?",
		newUser.Name, newUser.Family, newUser.UserName, newUser.Password)
	if yesOrNo() {
		savedata(newUser)

	} else {
		fmt.Println("No")

	}
	_ = mf.CharGetKey()
}

// --------------------------- Save to file --------------------------------
func savedata(nUser dataUser) {
	file, err := os.OpenFile("Userdata.json", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Can't create or open json file.....")
		return
	} else {
		defer file.Close()
		var cashData []dataUser
		stat, _ := file.Stat()
		if stat.Size() > 0 {

			redAll, err := io.ReadAll(file)
			if err != nil {
				fmt.Println("Read data for write agin have problem...", err)
				_ = mf.CharGetKey()
				return
			} else {
				err = json.Unmarshal(redAll, &cashData)
				if err != nil {
					fmt.Println("read convert jason have problem ....")
					_ = mf.CharGetKey()
					return
				}
			}
		}
		cashData = append(cashData, nUser)
		file.Truncate(0)
		file.Seek(0, 0)
		encode := json.NewEncoder(file)
		err = encode.Encode(cashData)
		if err != nil {
			fmt.Println("problem to save data ....")
			_ = mf.CharGetKey()
			return

		}

		fmt.Printf("Yes\n Data Saveing ....")

	}
}

// --------------------------- yes or no	-------------------------------
func yesOrNo() bool {
	for {
		read := mf.CharGetKey()
		if read == 'y' || read == 'Y' {
			return true
		} else {
			if read == 'n' || read == 'N' {
				return false
			}
		}
	}
}

// --------------------------Show ,Menu -------------------------------
func showMenu() int {
	mf.Cls()
	fmt.Println("...<<Manin manu forom User Managment system>>... ")
	fmt.Printf("1) Register\n2) Login\n3) Edit profile\n4) Delete Account\n5) Exit\nCommand: ")
	for {
		read := mf.CharGetKey()
		//fmt.Printf("read : %v type %T \n", read, read)
		if read == '1' || read == '2' || read == '3' || read == '4' || read == '5' {
			//fmt.Println(read)
			i, _ := strconv.Atoi(string(read))
			fmt.Printf("%v\n", i)
			return i
		}
	}
}
