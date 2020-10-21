import { ConflictException, InternalServerErrorException } from '@nestjs/common';
import { EntityRepository, Repository } from 'typeorm';
import { Create{{MODULE_NAME}}Dto } from './dto/create-{{MODULE_NAME_LOWER}}.dto';
import { Get{{MODULE_NAME}}FilterDto } from './dto/filter-{{MODULE_NAME_LOWER}}.dto';
import { {{MODULE_NAME}} } from './{{MODULE_NAME_LOWER}}.entity';

@EntityRepository({{MODULE_NAME}})
export class {{MODULE_NAME}}Repository extends Repository<{{MODULE_NAME}}> {
  /** New {{MODULE_NAME}} Registration */
  async create{{MODULE_NAME}}(create{{MODULE_NAME}}Dto: Create{{MODULE_NAME}}Dto): Promise<{{MODULE_NAME}}> {
    const {
    //   name,
    //   department,
    //   regno,
    //   gender,
    //   email_college,
    //   email_personal,
    //   instagram,
    //   twitter,
    //   isAlumni,
    //   isBoard{{MODULE_NAME}},
    //   linkedIn,
    //   github,
    } = create{{MODULE_NAME}}Dto;

    const {{MODULE_NAME_LOWER}} = this.create();
    // {{MODULE_NAME}}.name = name;
    // {{MODULE_NAME}}.department = department;
    // {{MODULE_NAME}}.regno = regno;
    // {{MODULE_NAME}}.gender = gender;
    // {{MODULE_NAME}}.email_college = email_college;
    // {{MODULE_NAME}}.email_personal = email_personal;
    // {{MODULE_NAME}}.instagram = instagram;
    // {{MODULE_NAME}}.twitter = twitter;
    // {{MODULE_NAME}}.isAlumni = isAlumni;
    // {{MODULE_NAME}}.isBoard{{MODULE_NAME}} = isBoard{{MODULE_NAME}};
    // {{MODULE_NAME}}.linkedIn = linkedIn;
    // {{MODULE_NAME}}.github = github;

    try {
      await {{MODULE_NAME_LOWER}}.save();
    } catch (error) {
      if (error.code === '23505') {
        throw new ConflictException();
      } else {
        throw new InternalServerErrorException();
      }
    }
    return {{MODULE_NAME_LOWER}};
  }

  /** filter {{MODULE_NAME_LOWER}}  */
  async get{{MODULE_NAME}}(filterDto: Get{{MODULE_NAME}}FilterDto): Promise<{{MODULE_NAME}}[]> {
    const { department } = filterDto;
    const query = this.createQueryBuilder('{{MODULE_NAME_LOWER}}');

    if (department) {
      query.andWhere('{{MODULE_NAME}}.department = :department', { department });
    }

    const result = await query.getMany();
    return result;
  }
}