import {
    Body,
    Controller,
    Delete,
    Get,
    Param,
    ParseIntPipe,
    Patch,
    Post,
    Query,
    UsePipes,
    ValidationPipe,
  } from '@nestjs/common';
  import { CreateUserDto } from './dto/create-user.dto';
  import { GetUserFilterDto } from './dto/filter-user.dto';
  import { UpdateUserDto } from './dto/update-user.dto';
  import { User } from './user.entity';
  import {UserService } from './user.service';
  
  @Controller('user')
  export class UserController {
    constructor(private userService: UserService) {}
  
    /** get list of all users */
    @Get()
    getUser(
      @Query(ValidationPipe) getUserFilterDto: GetUserFilterDto,
    ): Promise<User[]> {
      return this.userService.getUser(getUserFilterDto);
    }
  
    /** get list of all ids */
    @Get('/:id')
    getUserById(@Param('id', ParseIntPipe) id: number): Promise<User> {
      return this.userService.getUserById(id);
    }
  
    /** to create a new user entry */
    @Post()
    @UsePipes(ValidationPipe)
    createUser(@Body() createUserDto: CreateUserDto): Promise<User> {
      return this.userService.createUser(createUserDto);
    }
  
    /**  to remove a user */
    @Delete('/:id')
    removeUser(@Param('id', ParseIntPipe) id: number): Promise<void> {
      return this.userService.removeUser(id);
    }
  
    /** to update records */
    @Patch('/:id')
    updateUser(@Body() updateUserDto: UpdateUserDto): Promise<User> {
      return this.userService.updateUserDetails(updateUserDto);
    }
  }