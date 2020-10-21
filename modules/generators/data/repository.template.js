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
{{DESTRUCTURING}}
    } = create{{MODULE_NAME}}Dto;

    const {{MODULE_NAME_LOWER}} = this.create();
{{ASSIGNMENT}}
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
    //@NeedsManualIntervention
    const { department } = filterDto;
    const query = this.createQueryBuilder('{{MODULE_NAME_LOWER}}');

    if (department) {
      query.andWhere('{{MODULE_NAME}}.department = :department', { department });
    }

    const result = await query.getMany();
    return result;
  }
}