<<<<<<< HEAD
CREATe TABLE Student(
        Sno       CHAR(5) primary key,
  Sname  CHAR(20)  UNIQUE,
   Ssex      CHAR(1)  check (ssex in ('F','M')),
   Sage      INT,
   Sdept    CHAR(15)
)

CREATe TABLE course(
        Cno       CHAR(5) primary key,
  Cname  CHAR(20)  UNIQUE
)
CREATe TABLE course1(
        Cno       CHAR(5) ,
  Cname  CHAR(20)  UNIQUE,a
primary key (cno)
)

CREATE TABLE SC(
            Sno CHAR(5) ,
            Cno CHAR(5) , 
            Grade   int,
            Primary key (Sno, Cno),
     FOREIGN KEY (sno) REFERENCES student (sno) ,
     FOREIGN KEY (cno) REFERENCES course (cno) 
);

CREATE TABLE SC1(
            Sno CHAR(5) Primary key  ,
            Cno CHAR(5) Primary key , 
            Grade   int,
            
     FOREIGN KEY (sno) REFERENCES student (sno) ,
     FOREIGN KEY (cno) REFERENCES course (cno) 
)

CREATE TABLE SC1(
            Sno CHAR(5) ,
            Cno CHAR(5) , 
            Grade   int,
            Primary key (Sno, Cno),
     FOREIGN KEY (sno) REFERENCES student (sno) ,
     FOREIGN KEY (cno) REFERENCES course (cno) 
);

drop table Student2

alter table sc add COLUMN fcredit float
select * from student 
insert into student(sno,sname) values ('3301','zs')

insert student(sno,sname) values ('3302','LS');

insert student(sno,sname,sage) values ('33031',NULL,19),('33042','',20)

create table stu as select * from student 

create table stu1  select * from student 

create index wjlindexsname on student(sname)

create index wjlindexsnameage on student(sname,sage desc)
.


select * from student where sage=19 and sname='ws'
select * from student where sage=19 and sname=''
select * from student where sage=19 and sname=null

select * from student where  sname is not null

select * from student where  sname  not is null

select * from student where  not sname   is null

select sno,left(sno,2),sname from student order by sno desc 
select sno,substr(sno,3,3),sname from student order by sno desc 


select * from student order by rand() limit 1

select * from student;
select * from student  limit 2


select 1, sno,sage,sage+1,sname from student order by sno desc 


select * from student order by rand() limit 2,3

insert into course values (1,'Java'),(2,'DBMS')

insert into sc values (3301,1,80),(3302,2,90)
select * from sc

select sname from student where exists (select * from sc where sc.sno=student.sno and (cno=1 or cno=2) order by sno) order by sname desc 

select sname from student where not exists (select * from sc where sc.sno=student.sno and (cno=1 ) order by sno) order by sname desc 

select * from student where sno in (select sno from sc where cno=1 order by sno )

SELECT Sno,Sname,Sdept
     FROM Student S1
     WHERE EXISTS (SELECT *
           FROM Student S2
           WHERE S2.Sdept = S1.Sdept AND
                   S2.Sname = 'zs')

SELECT Sname
         FROM Student
         WHERE NOT EXISTS
            (SELECT *
              FROM Course
              WHERE NOT EXISTS
                  (SELECT *
                   FROM SC
                   WHERE Sno= Student.Sno
                      AND Cno= Course.Cno))

SELECT *
        FROM Student
        WHERE Sdept= 'CS'
        UNION
        SELECT *
        FROM Student
        WHERE Sage<=19

SELECT  DISTINCT  *
        FROM Student
        WHERE Sdept= 'CS'  OR  Sage<=19

        SELECT Sno
        FROM SC
        WHERE Cno='1'
        UNION
        SELECT Sno
        FROM SC
        WHERE Cno= '2'

SELECT  DISTINCT  Sno
        FROM SC
        WHERE Cno='1'  OR  Cno= '2'

       SELECT *
        FROM Student
        WHERE Sdept= 'CS' AND
              Sage<=19

       SELECT *
        FROM Student order by sname

SELECT Sno
        FROM SC
        WHERE Cno='2'

SELECT Sno
        FROM SC
        WHERE Cno='1' AND Sno IN
                               (SELECT Sno
                                FROM SC
                                WHERE Cno='2')

SELECT Sno
        FROM SC
        WHERE Cno='1' and cno='2'


SELECT   *
        FROM    Student
        WHERE Sdept= 'CS'

        UNION
        SELECT *
        FROM    Student
        WHERE Sage<=19
        ORDER BY  Sno


SELECT Sname
FROM Student, SC
WHERE Student.Sno=SC.Sno AND 
			SC.Cno= '1';


insert student(sno,sname,sage) values ('3305','ABC',19),('3306',"BCD",20)


select  sage ,sdept from student
select  distinct sage ,sdept from student


select distinct sage,distinct sdept from student


select distinct sdept from student

select * from student where not sname  in ('zs')
select * from student where  sname  like 'A%b' escape 'A'
select * from student where not sname  like ''

select * from student where  sname  not like 'zs'

select * from student where not sname  is  null
select * from student where  sname is not null
select * from student where  sname = null

select * from student order by sname desc
update  student set ssex='M'  where sno >'3303'

select *,count(*) as num from student
select count(*) as num from student
select  max(sage),min(sage),sum(sage),avg(sage),count(DISTINCT sage) from student

select ssex,count(*) from student group by ssex  having count(*) >2


select ssex,sname,count(*) from stu where ssex='F' group by ssex ,sname having count(*) >1

select * from sc
SELECT Sno
     FROM  SC
     GROUP BY Sno
     HAVING  COUNT(*) >1



SELECT  Student.* ,  SC.* ,  course.*
         FROM     Student, SC WHERE  Student.Sno = SC.Sno and 
course.cno=sc.cno

select * from course


SELECT  s.sno,sname,course.cno,cname,grade
         FROM     Student as s , SC,course WHERE  s.Sno = SC.Sno and 
course.cno=sc.cno

SELECT  Student.Sno,Sname,Ssex,Sage,Sdept,Cno,Grade
 FROM   SC  right join  Student
 on  Student.Sno = SC.Sno
=======
CREATe TABLE Student(
        Sno       CHAR(5) primary key,
  Sname  CHAR(20)  UNIQUE,
   Ssex      CHAR(1)  check (ssex in ('F','M')),
   Sage      INT,
   Sdept    CHAR(15)
)

CREATe TABLE course(
        Cno       CHAR(5) primary key,
  Cname  CHAR(20)  UNIQUE
)
CREATe TABLE course1(
        Cno       CHAR(5) ,
  Cname  CHAR(20)  UNIQUE,a
primary key (cno)
)

CREATE TABLE SC(
            Sno CHAR(5) ,
            Cno CHAR(5) , 
            Grade   int,
            Primary key (Sno, Cno),
     FOREIGN KEY (sno) REFERENCES student (sno) ,
     FOREIGN KEY (cno) REFERENCES course (cno) 
);

CREATE TABLE SC1(
            Sno CHAR(5) Primary key  ,
            Cno CHAR(5) Primary key , 
            Grade   int,
            
     FOREIGN KEY (sno) REFERENCES student (sno) ,
     FOREIGN KEY (cno) REFERENCES course (cno) 
)

CREATE TABLE SC1(
            Sno CHAR(5) ,
            Cno CHAR(5) , 
            Grade   int,
            Primary key (Sno, Cno),
     FOREIGN KEY (sno) REFERENCES student (sno) ,
     FOREIGN KEY (cno) REFERENCES course (cno) 
);

drop table Student2

alter table sc add COLUMN fcredit float
select * from student 
insert into student(sno,sname) values ('3301','zs')

insert student(sno,sname) values ('3302','LS');

insert student(sno,sname,sage) values ('33031',NULL,19),('33042','',20)

create table stu as select * from student 

create table stu1  select * from student 

create index wjlindexsname on student(sname)

create index wjlindexsnameage on student(sname,sage desc)
.


select * from student where sage=19 and sname='ws'
select * from student where sage=19 and sname=''
select * from student where sage=19 and sname=null

select * from student where  sname is not null

select * from student where  sname  not is null

select * from student where  not sname   is null

select sno,left(sno,2),sname from student order by sno desc 
select sno,substr(sno,3,3),sname from student order by sno desc 


select * from student order by rand() limit 1

select * from student;
select * from student  limit 2


select 1, sno,sage,sage+1,sname from student order by sno desc 


select * from student order by rand() limit 2,3

insert into course values (1,'Java'),(2,'DBMS')

insert into sc values (3301,1,80),(3302,2,90)
select * from sc

select sname from student where exists (select * from sc where sc.sno=student.sno and (cno=1 or cno=2) order by sno) order by sname desc 

select sname from student where not exists (select * from sc where sc.sno=student.sno and (cno=1 ) order by sno) order by sname desc 

select * from student where sno in (select sno from sc where cno=1 order by sno )

SELECT Sno,Sname,Sdept
     FROM Student S1
     WHERE EXISTS (SELECT *
           FROM Student S2
           WHERE S2.Sdept = S1.Sdept AND
                   S2.Sname = 'zs')

SELECT Sname
         FROM Student
         WHERE NOT EXISTS
            (SELECT *
              FROM Course
              WHERE NOT EXISTS
                  (SELECT *
                   FROM SC
                   WHERE Sno= Student.Sno
                      AND Cno= Course.Cno))

SELECT *
        FROM Student
        WHERE Sdept= 'CS'
        UNION
        SELECT *
        FROM Student
        WHERE Sage<=19

SELECT  DISTINCT  *
        FROM Student
        WHERE Sdept= 'CS'  OR  Sage<=19

        SELECT Sno
        FROM SC
        WHERE Cno='1'
        UNION
        SELECT Sno
        FROM SC
        WHERE Cno= '2'

SELECT  DISTINCT  Sno
        FROM SC
        WHERE Cno='1'  OR  Cno= '2'

       SELECT *
        FROM Student
        WHERE Sdept= 'CS' AND
              Sage<=19

       SELECT *
        FROM Student order by sname

SELECT Sno
        FROM SC
        WHERE Cno='2'

SELECT Sno
        FROM SC
        WHERE Cno='1' AND Sno IN
                               (SELECT Sno
                                FROM SC
                                WHERE Cno='2')

SELECT Sno
        FROM SC
        WHERE Cno='1' and cno='2'


SELECT   *
        FROM    Student
        WHERE Sdept= 'CS'

        UNION
        SELECT *
        FROM    Student
        WHERE Sage<=19
        ORDER BY  Sno


SELECT Sname
FROM Student, SC
WHERE Student.Sno=SC.Sno AND 
			SC.Cno= '1';


insert student(sno,sname,sage) values ('3305','ABC',19),('3306',"BCD",20)


select  sage ,sdept from student
select  distinct sage ,sdept from student


select distinct sage,distinct sdept from student


select distinct sdept from student

select * from student where not sname  in ('zs')
select * from student where  sname  like 'A%b' escape 'A'
select * from student where not sname  like ''

select * from student where  sname  not like 'zs'

select * from student where not sname  is  null
select * from student where  sname is not null
select * from student where  sname = null

select * from student order by sname desc
update  student set ssex='M'  where sno >'3303'

select *,count(*) as num from student
select count(*) as num from student
select  max(sage),min(sage),sum(sage),avg(sage),count(DISTINCT sage) from student

select ssex,count(*) from student group by ssex  having count(*) >2


select ssex,sname,count(*) from stu where ssex='F' group by ssex ,sname having count(*) >1

select * from sc
SELECT Sno
     FROM  SC
     GROUP BY Sno
     HAVING  COUNT(*) >1



SELECT  Student.* ,  SC.* ,  course.*
         FROM     Student, SC WHERE  Student.Sno = SC.Sno and 
course.cno=sc.cno

select * from course


SELECT  s.sno,sname,course.cno,cname,grade
         FROM     Student as s , SC,course WHERE  s.Sno = SC.Sno and 
course.cno=sc.cno

SELECT  Student.Sno,Sname,Ssex,Sage,Sdept,Cno,Grade
 FROM   SC  right join  Student
 on  Student.Sno = SC.Sno
>>>>>>> main_1
