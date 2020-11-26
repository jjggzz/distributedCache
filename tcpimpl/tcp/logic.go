package tcp

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net"
	"strings"
)

func (s *Server) logic(conn net.Conn) {
	// 关闭资源
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	reader := bufio.NewReader(conn)
	// 获取操作命令
	op, err := reader.ReadString(' ')
	if err != nil {
		if err == io.EOF {
			return
		}
		log.Println(err)
	}
	op = strings.TrimSpace(op)
	// 处理
	if op == "get" {
		bytes, err := s.get(reader)
		response(conn, bytes, err)
	} else if op == "set" {
		err := s.set(reader)
		response(conn, []byte("操作成功"), err)
	} else if op == "del" {
		err := s.del(reader)
		response(conn, []byte("操作成功"), err)
	} else {
		bytes, err := s.stat()
		response(conn, bytes, err)
	}

}

func (s *Server) get(r *bufio.Reader) ([]byte, error) {
	//获取key
	key, err := r.ReadString(' ')
	if err != nil {
		if err == io.EOF {
			return nil, errors.New("key为空")
		}
		log.Println(err)
	}
	return s.Get(strings.TrimSpace(key))
}

func (s *Server) set(r *bufio.Reader) error {
	//获取key
	key, err := r.ReadString(' ')
	if err != nil {
		if err == io.EOF {
			return errors.New("key为空")
		}
		log.Println(err)
	}
	//获取剩下的value
	value, err := r.ReadString(' ')
	if err != nil {
		if err == io.EOF {
			return errors.New("value为空")
		}
		log.Println(err)
	}
	return s.Set(strings.TrimSpace(key), []byte(strings.TrimSpace(value)))
}

func (s *Server) del(r *bufio.Reader) error {
	//获取key
	key, err := r.ReadString(' ')
	if err != nil {
		if err == io.EOF {
			return errors.New("key为空")
		}
		log.Println(err)
	}
	return s.Del(strings.TrimSpace(key))
}

func (s *Server) stat() ([]byte, error) {
	stat := s.GetStat()
	bytes, err := json.Marshal(stat)
	if err != nil {
		return nil, errors.New("获取状态失败")
	}
	return bytes, err
}

func response(conn net.Conn, data []byte, err error) {
	if err != nil {
		_, _ = conn.Write([]byte(err.Error()))
		return
	}
	_, _ = conn.Write(data)
}
