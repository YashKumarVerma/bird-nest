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
  import { Create{{MODULE_NAME}}Dto } from './dto/create-{{MODULE_NAME_LOWER}}.dto';
  import { Get{{MODULE_NAME}}FilterDto } from './dto/filter-{{MODULE_NAME_LOWER}}.dto';
  import { Update{{MODULE_NAME}}Dto } from './dto/update-{{MODULE_NAME_LOWER}}.dto';
  import { {{MODULE_NAME}} } from './{{MODULE_NAME_LOWER}}.entity';
  import {{{MODULE_NAME}}Service } from './{{MODULE_NAME_LOWER}}.service';
  
  @Controller('{{MODULE_NAME_LOWER}}')
  export class {{MODULE_NAME}}Controller {
    constructor(private {{MODULE_NAME_LOWER}}Service: {{MODULE_NAME}}Service) {}
  
    /** get list of all {{MODULE_NAME_LOWER}}s */
    @Get()
    get{{MODULE_NAME}}(
      @Query(ValidationPipe) get{{MODULE_NAME}}FilterDto: Get{{MODULE_NAME}}FilterDto,
    ): Promise<{{MODULE_NAME}}[]> {
      return this.{{MODULE_NAME_LOWER}}Service.get{{MODULE_NAME}}(get{{MODULE_NAME}}FilterDto);
    }
  
    /** get list of all ids */
    @Get('/:id')
    get{{MODULE_NAME}}ById(@Param('id', ParseIntPipe) id: number): Promise<{{MODULE_NAME}}> {
      return this.{{MODULE_NAME_LOWER}}Service.get{{MODULE_NAME}}ById(id);
    }
  
    /** to create a new {{MODULE_NAME_LOWER}} entry */
    @Post()
    @UsePipes(ValidationPipe)
    create{{MODULE_NAME}}(@Body() create{{MODULE_NAME}}Dto: Create{{MODULE_NAME}}Dto): Promise<{{MODULE_NAME}}> {
      return this.{{MODULE_NAME_LOWER}}Service.create{{MODULE_NAME}}(create{{MODULE_NAME}}Dto);
    }
  
    /**  to remove a {{MODULE_NAME_LOWER}} */
    @Delete('/:id')
    remove{{MODULE_NAME}}(@Param('id', ParseIntPipe) id: number): Promise<void> {
      return this.{{MODULE_NAME_LOWER}}Service.remove{{MODULE_NAME}}(id);
    }
  
    /** to update records */
    @Patch('/:id')
    update{{MODULE_NAME}}(@Body() update{{MODULE_NAME}}Dto: Update{{MODULE_NAME}}Dto): Promise<{{MODULE_NAME}}> {
      return this.{{MODULE_NAME_LOWER}}Service.update{{MODULE_NAME}}Details(update{{MODULE_NAME}}Dto);
    }
  }