# 作业

我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

dao层应该wrap这个error，抛给上层来处理，因为dao层并不能决定怎么处理这个错误，应该留给业务层来解决。

代码中我使用一个HTTP服务来演示，dao目录下为dao层，handler目录下为业务层，在dao层的错误都通过errors包来wrap起来，在handlers中的代码来处理。

# 学习笔记

1. https://www.digitalocean.com/community/tutorials/how-to-create-a-new-user-and-grant-permissions-in-mysql

2. If you started mysql using mysql -u root -p
   
   Try ALTER USER 'root'@'localhost' IDENTIFIED BY 'MyNewPass';
   
   Source: http://dev.mysql.com/doc/refman/5.7/en/resetting-permissions.html
   
3. https://dzone.com/articles/docker-for-mac-mysql-setup

4. https://blog.golang.org/go1.13-errors

5. https://dev.mysql.com/doc/mysql-installation-excerpt/8.0/en/docker-mysql-getting-started.html

6. https://severalnines.com/database-blog/mysql-docker-containers-understanding-basics

7. https://github.com/gustavocd/dao-pattern-in-go

8. https://stackoverflow.com/questions/43748751/golang-dtos-entities-and-mapping

9. https://medium.com/capital-one-tech/doing-well-by-doing-bad-writing-bad-code-with-go-part-2-e270d305c9f7

10. https://www.alexedwards.net/blog/organising-database-access

11. https://scene-si.org/2018/01/25/go-tips-and-tricks-almost-everything-about-imports/

12. https://blog.golang.org/using-go-modules

