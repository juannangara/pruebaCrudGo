use school

CREATE TABLE IF NOT EXISTS `school`.`Course` (
  `courseID` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(15) NOT NULL,
  `crationDate` date NOT NULL,
  `description` VARCHAR(50) NOT NULL,
  PRIMARY KEY (`courseID`))
 ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `school`.`Grade` (
  `gradeID` INT NOT NULL AUTO_INCREMENT,
  `grade` INT NOT NULL,
  `crationDate` date NOT NULL,
  `courseID` INT NOT NULL,
 PRIMARY KEY (`gradeID`),
 CONSTRAINT `fk_course_grade`
    FOREIGN KEY (`courseID`)
    REFERENCES  course(`courseID`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
   ENGINE = InnoDB;
  
CREATE TABLE IF NOT EXISTS `school`.`Student` (
  `studentID` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(15) NOT NULL,
  `crationDate` date NOT NULL,
  PRIMARY KEY (`studentID`))
 ENGINE = InnoDB;;
 
 
CREATE TABLE IF NOT EXISTS `school`.`GradesPerStudent` (
  `studentID` INT NOT NULL ,
  `gradeID` INT NOT NULL,
  CONSTRAINT `fk_student`
    FOREIGN KEY (`studentID`)
    REFERENCES `school`.`Student` (`studentID`),
  CONSTRAINT `fk_grade`
    FOREIGN KEY (`gradeID`)
    REFERENCES `school`.`Grade` (`gradeID`))
   ENGINE = InnoDB;
    
CREATE TABLE IF NOT EXISTS `school`.`StudentPerCourse` (
  `studentID` INT NOT NULL ,
  `courseID` INT NOT NULL,
  CONSTRAINT `fk_student_course`
    FOREIGN KEY (`studentID`)
    REFERENCES `school`.`Student` (`studentID`),
  CONSTRAINT `fk_course_ student`
    FOREIGN KEY (`courseID`)
    REFERENCES `school`.`Course` (`courseID`))
    ENGINE = InnoDB;