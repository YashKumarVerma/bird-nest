import { BaseEntity, Column, Entity, PrimaryGeneratedColumn } from 'typeorm';
import {  } from './User.enum';

@Entity()
export class User extends BaseEntity {

@PrimaryGeneratedColumn()
id : number;

@Column({unique : true})
name : string;

@Column()
age : number;

@Column({default : 0 })
money : number;

@Column()
gender : string;

@Column()
devices : string;

@Column({default : "in" ,length : "2" })
country : string;

@Column({unique : true})
github : string;

@Column({unique : true})
email : string;

@Column({unique : true})
instagram : string;

}