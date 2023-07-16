# go-OracleTranscoder
当修改NLS_LANG也不能oracle乱码时，可以使用此工具

原理是把oracle查询的中文结果转换为16进制，然后在本地解码输出。
中文作为入参时，做反向操作。