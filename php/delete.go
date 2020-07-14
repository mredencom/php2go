package php

// https://www.php.net/manual/zh/function.delete.php
// delete()
func Delete(filename string) bool {
	return Unlink(filename)
}
