import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { User } from './user.entity';
import { CreateUserDto } from './dto/create-user.dto';
import { UserRepository } from './user.repository';
import { GetUserFilterDto } from './dto/filter-user.dto';
import { UpdateUserDto } from './dto/update-user.dto';

@Injectable()
export class UserService {
  /** injecting dependencies */
  constructor(
    @InjectRepository(UserRepository)
    private UserRepository: UserRepository,
  ) {}

  /** create new user */
  async createUser(createUserDto: CreateUserDto): Promise<User> {
    return await this.userRepository.createUser(createUserDto);
  }

  /** to remove a user */
  async removeUser(id: number): Promise<void> {
    const result = await this.userRepository.delete(id);

    if (result.affected === 0) {
      throw new NotFoundException(`User with ${id} does not exist`);
    }
  }

  /** to retrieve user list */
  async getUser(filterDto: GetUserFilterDto): Promise<User[]> {
    return await this.userRepository.getUser(filterDto);
  }

  /** get user by id */
  async getUserById(id: number): Promise<User> {
    const found = await this.userRepository.findOne(id);
    if (!found) {
      throw new NotFoundException(`User with ${id} not created yet`);
    }
    return found;
  }

  /** update User details */
  async updateUserDetails(updateUserDto: UpdateUserDto): Promise<User> {
    return await this.userRepository.save({ ...updateUserDto });
  }
}