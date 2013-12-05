package os_test

import (
	"errors"
	"os"
	"testing"
	"time"
)

//EWINDOWS - APPLICATION_ERROR:        "not supported by windows",

// TODO(brainman): fix all needed for os
// func Getppid() (ppid int) { return -1 }

// func Fchdir(fd Handle) (err error)                        { return EWINDOWS }
// func Link(oldpath, newpath string) (err error)            { return EWINDOWS }
// func Symlink(path, link string) (err error)               { return EWINDOWS }
// func Readlink(path string, buf []byte) (n int, err error) { return 0, EWINDOWS }

// func Fchmod(fd Handle, mode uint32) (err error)        { return EWINDOWS }
// func Chown(path string, uid int, gid int) (err error)  { return EWINDOWS }
// func Lchown(path string, uid int, gid int) (err error) { return EWINDOWS }
// func Fchown(fd Handle, uid int, gid int) (err error)   { return EWINDOWS }

// func Getuid() (uid int)                  { return -1 }
// func Geteuid() (euid int)                { return -1 }
// func Getgid() (gid int)                  { return -1 }
// func Getegid() (egid int)                { return -1 }
// func Getgroups() (gids []int, err error) { return nil, EWINDOWS }

func TestChdir_存在するディレクトリに移動 (t *testing.T) {

	current := printCurrentDir(t)

	defer os.Chdir(current)

	err := os.Chdir("c:/windows")

	if err != nil {
		t.Fatal(err)
	}

	printCurrentDir(t)

}

func TestChdir_存在しないディレクトリに移動 (t *testing.T) {

	current := printCurrentDir(t)

	defer os.Chdir(current)

	err := os.Chdir("z:/")

	if err != nil {
		t.Log(err)
	} else {
		t.Fatal("ここに到達してはいけません。")
	}

}

func TestChmod(t *testing.T) {

	const filePath = "./sample.txt"

	err := os.Chmod(filePath, 0777)

	if err != nil {
		t.Fatal(err)
	}
	
	printFileInfo(t, filePath)

	err = os.Chmod(filePath, 0400)

	if err != nil {
		t.Fatal(err)
	}
	
	printFileInfo(t, filePath)	

}

func TestChown(t *testing.T) {

	const filePath = "./sample.txt"

	err := os.Chown(filePath, 1, 1)

	t.Skip("windowsでは未対応です。", err)

}

func TestChtimes(t *testing.T) {

	const filePath = "./sample.txt"

	printFileInfo(t, filePath)

	atime := time.Now().Add(1 * time.Hour)
	mtime := time.Now().Add(2 * time.Hour)

	err := os.Chtimes(filePath, atime, mtime)

	if err != nil {
		t.Fatal(err)
	}

	// TODO: I don't know how to get access time
	printFileInfo(t, filePath)

}

// 他のテストに影響があるかもしれないのでコメントアウト
// func TestClearenv(t *testing.T) {

// 	printCurrentEnv(t)

// 	os.Clearenv()

// 	printCurrentEnv(t)

// }

func TestEnviron(t *testing.T) {

	printCurrentEnv(t)

}

// テストが終了してしまうのでコメントアウト
// func TestExit(t *testing.T) {

// 	os.Exit(0)

// }

func TestExpand(t *testing.T) {

	str := os.Expand("GOROOTは $GOROOT です。", func(s string) string {
			return "###" + s + "###"
		})

	t.Log(str)

}

func TestExpandEnv_環境変数定義あり(t *testing.T) {

	t.Log(os.ExpandEnv("GOROOTは $GOROOT です。"))
	t.Log(os.ExpandEnv("GOROOTは ${GOROOT} です。"))

}

func TestExpandEnv_環境変数定義なし(t *testing.T) {

	t.Log(os.ExpandEnv("GOROOTは $GOROOT1 です。"))

}

func TestGetegid(t *testing.T) {

	t.Skip("windowsでは未対応です。", os.Getegid())

}

func TestGetenv(t *testing.T) {

	t.Log("GOROOT", os.Getenv("GOROOT"))

}

func TestGeteuid(t *testing.T) {

	t.Skip("windowsでは未対応です。", os.Geteuid())

}

func TestGetgid(t *testing.T) {

	t.Skip("windowsでは未対応です。", os.Getgid())

}

func TestGetGroups(t *testing.T) {

	groups, err := os.Getgroups()

	t.Skip("windowsでは未対応です。", groups, err)

}

func TestGetpagesize(t *testing.T) {

	t.Log(os.Getpagesize())

}

func TestGetpid(t *testing.T) {

	t.Log(os.Getpid())

}

func TestGetppid(t *testing.T) {

	t.Skip("windowsでは未対応です。", os.Getppid())

}

func TestGetuid(t *testing.T) {

	t.Skip("windowsでは未対応です。", os.Getuid())

}

func TestGetwd(t *testing.T) {

	printCurrentDir(t)

}

func TestHostname(t *testing.T) {

	h, err := os.Hostname()

	if err != nil {
		t.Fatal(err)
	}

	t.Log(h)

}

func TestIsExist(t *testing.T) {

	// error が ErrExist かどうか判定する
	t.Log(os.ErrExist, os.IsExist(os.ErrExist))
	t.Log(os.ErrPermission, os.IsExist(os.ErrPermission))

	err := errors.New("自分で定義したエラー")
	t.Log(err, os.IsExist(err))

}

func TestIsNotExist(t *testing.T) {

	// error が ErrNotExist かどうか判定する
	t.Log(os.ErrNotExist, os.IsNotExist(os.ErrNotExist))
	t.Log(os.ErrPermission, os.IsExist(os.ErrPermission))

	err := errors.New("自分で定義したエラー")
	t.Log(err, os.IsNotExist(err))

}

func TestIsPathSeparator(t *testing.T) {

	t.Log("/", os.IsPathSeparator('/'))
	t.Log("\\", os.IsPathSeparator('\\'))
	t.Log(":", os.IsPathSeparator(':'))
	t.Log("@", os.IsPathSeparator('@'))

}

func TestIsPermission(t *testing.T) {

	// error が ErrPermission かどうか判定する
	t.Log(os.ErrPermission, os.IsPermission(os.ErrPermission))
	t.Log(os.ErrExist, os.IsPermission(os.ErrExist))

	err := errors.New("自分で定義したエラー")
	t.Log(err, os.IsPermission(err))

}

func TestLchown(t *testing.T) {

	const filePath = "./sample.txt"

	err := os.Lchown(filePath, 1, 1)

	t.Skip("windowsでは未対応です。", err)

}

func TestLink(t *testing.T) {

	err := os.Link("./sample.txt", "./sample2.txt")

	t.Skip("windowsでは未対応です。", err)

}

func TestMkdir(t *testing.T) {

	dir := createDirectory(t, "gotest")

	t.Log(dir)
	
	printFileInfo(t, dir)

	defer os.Remove(dir)

}

func TestMkdir_すでに存在するディレクトリを作成しようとした場合(t *testing.T) {

	dir := createDirectory(t, "gotest")

	t.Log(dir)

	defer os.Remove(dir)

	err := os.Mkdir(dir, 0777)

	if err != nil {
		t.Log(err)
	} else {
		t.Fatal("ここには到達しないはずです。")
	}

}

func TestMkdirAll(t *testing.T) {

	dir := createDirectoryAll(t, "gotest/gotest2/gotest3")

	t.Log(dir)

	printFileInfo(t, dir)

	defer os.Remove(dir)

}

func TestMkdirAll_すでに存在するディレクトリを作成しようとした場合(t *testing.T) {

	dir := createDirectoryAll(t, "gotest/gotest2/gotest3")

	t.Log(dir)

	defer os.Remove(dir)

	err := os.Mkdir(dir, 0777)

	// ドキュメントではエラーにならないはずだが
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log("ok")
	}

}

func TestReadlink(t *testing.T) {

	l, err := os.Readlink("D:/study/html5")

	if err != nil {
		t.Skip(err)
	}

	t.Log(l)

}

func TestRemove_存在するディレクトリ(t *testing.T) {

	dir := createDirectory(t, "gotest2")

	err := os.Remove(dir)

	if err != nil {
		t.Fatal(err)
	}

	t.Log("removed", dir)

}

func TestRemove_存在しないディレクトリ(t *testing.T) {

	err := os.Remove("w:/hoge")

	if err != nil {
		t.Log(err)
	} else {
		t.Fatal("ここには到達しないはずです。")
	}

}

func TestRemove_ディレクトリ指定なし(t *testing.T) {

	err := os.Remove("")

	if err != nil {
		t.Log(err)
	} else {
		t.Fatal("ここには到達しないはずです。")
	}

}

func TestRemoveAll(t *testing.T) {

	dir := createDirectory(t, "gotest/gotest2/gotest3")

	err := os.RemoveAll("gotest")

	if err != nil {
		t.Fatal(err)
	}

	t.Log("removed", dir)

}

func printCurrentDir(t *testing.T) (current string) {
	current, err := os.Getwd()

	if err != nil {
		t.Fatal(err)
	} else {
		t.Log("current", current)
	}

	return
}

func printFileInfo(t *testing.T, filePath string) {
	f, err := os.Open(filePath)

	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	fi, err := f.Stat()

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s, %d, %v, %v, %v", fi.Name(), fi.Size(), fi.Mode(), fi.ModTime(), fi.IsDir())
}

func printCurrentEnv(t *testing.T) {
	for _, env := range os.Environ() {
		t.Log(env)
	}
}

func createDirectory(t *testing.T, dirPath string) string {
	dir := os.TempDir() + "/" + dirPath

	err := os.Mkdir(dir, 0777)

	if err != nil {
		t.Fatal(err)
	}

	return dir
}

func createDirectoryAll(t *testing.T, dirPath string) string {
	dir := os.TempDir() + "/" + dirPath

	err := os.MkdirAll(dir, 0777)

	if err != nil {
		t.Fatal(err)
	}

	return dir
}
