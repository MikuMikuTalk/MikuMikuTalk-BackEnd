package etcd

import (
	"context"

	"im_server/core"

	"github.com/zeromicro/go-zero/core/logx"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdStorage struct {
	client *clientv3.Client
}

func NewEtcdStorage(addr string) *EtcdStorage {
	client := core.InitEtcd(addr)
	return &EtcdStorage{
		client: client,
	}
}

func (s *EtcdStorage) Put(key, value string) error {
	_, err := s.client.Put(context.Background(), key, value)
	if err != nil {
		logx.Errorf("Failed to put key %s: %v", key, err)
		return err
	}
	return err
}

// Get 获取配置
func (s *EtcdStorage) Get(key string) (string, error) {
	resp, err := s.client.Get(context.Background(), key)
	if err != nil {
		logx.Errorf("Failed to get key %s: %v", key, err)
		return "", err
	}
	if len(resp.Kvs) == 0 {
		return "", nil // Key 不存在
	}
	return string(resp.Kvs[0].Value), nil
}

// Delete 删除配置
func (s *EtcdStorage) Delete(key string) error {
	_, err := s.client.Delete(context.Background(), key)
	if err != nil {
		logx.Errorf("Failed to delete key %s: %v", key, err)
	}
	return err
}
