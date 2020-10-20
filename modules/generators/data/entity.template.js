import { BaseEntity, Column, Entity, PrimaryGeneratedColumn } from 'typeorm';
import {  } from './{{MODULE_NAME}}.enum';

@Entity()
export class {{MODULE_NAME}} extends BaseEntity {
 {{MAIN_WORKSPACE}}
}