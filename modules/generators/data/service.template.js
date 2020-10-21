import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { {{MODULE_NAME}} } from './{{MODULE_NAME_LOWER}}.entity';
import { Create{{MODULE_NAME}}Dto } from './dto/create-{{MODULE_NAME_LOWER}}.dto';
import { {{MODULE_NAME}}Repository } from './{{MODULE_NAME_LOWER}}.repository';
import { Get{{MODULE_NAME}}FilterDto } from './dto/filter-{{MODULE_NAME_LOWER}}.dto';
import { Update{{MODULE_NAME}}Dto } from './dto/update-{{MODULE_NAME_LOWER}}.dto';

@Injectable()
export class {{MODULE_NAME}}Service {
  /** injecting dependencies */
  constructor(
    @InjectRepository({{MODULE_NAME}}Repository)
    private {{MODULE_NAME}}Repository: {{MODULE_NAME}}Repository,
  ) {}

  /** create new {{MODULE_NAME_LOWER}} */
  async create{{MODULE_NAME}}(create{{MODULE_NAME}}Dto: Create{{MODULE_NAME}}Dto): Promise<{{MODULE_NAME}}> {
    return await this.{{MODULE_NAME_LOWER}}Repository.create{{MODULE_NAME}}(create{{MODULE_NAME}}Dto);
  }

  /** to remove a {{MODULE_NAME_LOWER}} */
  async remove{{MODULE_NAME}}(id: number): Promise<void> {
    const result = await this.{{MODULE_NAME_LOWER}}Repository.delete(id);

    if (result.affected === 0) {
      throw new NotFoundException(`{{MODULE_NAME}} with ${id} does not exist`);
    }
  }

  /** to retrieve {{MODULE_NAME_LOWER}} list */
  async get{{MODULE_NAME}}(filterDto: Get{{MODULE_NAME}}FilterDto): Promise<{{MODULE_NAME}}[]> {
    return await this.{{MODULE_NAME_LOWER}}Repository.get{{MODULE_NAME}}(filterDto);
  }

  /** get {{MODULE_NAME_LOWER}} by id */
  async get{{MODULE_NAME}}ById(id: number): Promise<{{MODULE_NAME}}> {
    const found = await this.{{MODULE_NAME_LOWER}}Repository.findOne(id);
    if (!found) {
      throw new NotFoundException(`User with ${id} not created yet`);
    }
    return found;
  }

  /** update {{MODULE_NAME}} details */
  async update{{MODULE_NAME}}Details(update{{MODULE_NAME}}Dto: Update{{MODULE_NAME}}Dto): Promise<{{MODULE_NAME}}> {
    return await this.{{MODULE_NAME_LOWER}}Repository.save({ ...update{{MODULE_NAME}}Dto });
  }
}