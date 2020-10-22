import { ConflictException, InternalServerErrorException } from '@nestjs/common';
import { EntityRepository, Repository } from 'typeorm';
import { CreateUserDto } from './dto/create-user.dto';
import { GetUserFilterDto } from './dto/filter-user.dto';
import { User } from './user.entity';

@EntityRepository(User)
export class UserRepository extends Repository<User> {
  /** New User Registration */
  async createUser(createUserDto: CreateUserDto): Promise<User> {
    const {
	id,
	name,
	age,
	money,
	gender,
	devices,
	country,
	github,
	email,
	instagram,

    } = createUserDto;

    const user = this.create();
	user.id = id;
	user.name = name;
	user.age = age;
	user.money = money;
	user.gender = gender;
	user.devices = devices;
	user.country = country;
	user.github = github;
	user.email = email;
	user.instagram = instagram;

    try {
      await user.save();
    } catch (error) {
      if (error.code === '23505') {
        throw new ConflictException();
      } else {
        throw new InternalServerErrorException();
      }
    }
    return user;
  }

  /** filter user  */
  async getUser(filterDto: GetUserFilterDto): Promise<User[]> {
    //@NeedsManualIntervention
    const { } = filterDto;
    const query = this.createQueryBuilder('user');

    // if (department) {
    //   query.andWhere('User.department = :department', { department });
    // }

    const result = await query.getMany();
    return result;
  }
}