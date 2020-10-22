import { } from '../user.enum';
//@NeedsManualIntervention
import { IsAlpha, IsEmail, IsIn, IsNotEmpty, Length } from 'class-validator';
import { ApiProperty } from '@nestjs/swagger';

export class CreateUserDto {
	@ApiProperty({ description : 'something'})
	@IsNotEmpty()
	id: int;

	@ApiProperty({ description : 'something'})
	@IsNotEmpty()
	@IsAlpha()
	name: string;

	@ApiProperty({ description : 'something'})
	@IsNotEmpty()
	age: int;

	@ApiProperty({ description : 'something'})
	@IsNotEmpty()
	money: int;

	@ApiProperty({ description : 'something'})
	@IsNotEmpty()
	@IsAlpha()
	gender: string;

	@ApiProperty({ description : 'something'})
	@IsNotEmpty()
	@IsAlpha()
	devices: string;

	@ApiProperty({ description : 'something'})
	@Length(2)
	@IsNotEmpty()
	@IsAlpha()
	country: string;

	@ApiProperty({ description : 'something'})
	@IsNotEmpty()
	@IsAlpha()
	github: string;

	@ApiProperty({ description : 'something'})
	@IsNotEmpty()
	@IsAlpha()
	email: string;

	@ApiProperty({ description : 'something'})
	@IsNotEmpty()
	@IsAlpha()
	instagram: string;


}