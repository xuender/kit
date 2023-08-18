// Package cfg 读取可加密配置信息.
// 对配置文件中占位符 AES[明文] 或 DES[明文] 加密，并返回明文.
// 原有占位符被替换成 AES(密文) 或 DES(密文).
// 主要方法 cfg.New(password).Bytes(path), cfg.New(password).Reader(path), cfg.New(password).String(path).
package cfg
