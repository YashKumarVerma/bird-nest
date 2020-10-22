import { } from '../user.enum';
import { IsEmail, IsIn, IsNumber, IsOptional, Length } from 'class-validator';
import { ApiProperty } from '@nestjs/swagger';

export class UpdateUserDto {
	@ApiProperty({
		description: 'something',
		required: true
	})
	@IsNumber()
	id: number;

	@ApiProperty({
		description: 'something',
		required: false
	})
	@IsOptional()
	name : string;

	@ApiProperty({
		description: 'something',
		required: false
	})
	@IsOptional()
	age : number;

	@ApiProperty({
		description: 'something',
		required: false
	})
	@IsOptional()
	money : number;

	@ApiProperty({
		description: 'something',
		required: false
	})
	@IsOptional()
	gender : string;

	@ApiProperty({
		description: 'something',
		required: false
	})
	@IsOptional()
	devices : string;

	@ApiProperty({
		description: 'something',
		required: false
	})
	@IsOptional()
	@Length(2)
	country : string;

	@ApiProperty({
		description: 'something',
		required: false
	})
	@IsOptional()
	github : string;

	@ApiProperty({
		description: 'something',
		required: false
	})
	@IsOptional()
	email : string;

	@ApiProperty({
		description: 'something',
		required: false
	})
	@IsOptional()
	instagram : string;

    
}