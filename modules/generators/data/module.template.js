import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { {{MODULE_NAME}}Controller } from './{{MODULE_NAME_LOWER}}.controller';
import { {{MODULE_NAME}}Repository } from './{{MODULE_NAME_LOWER}}.repository';
import { {{MODULE_NAME}}Service } from './{{MODULE_NAME_LOWER}}.service';

@Module({
  imports: [TypeOrmModule.forFeature([{{MODULE_NAME}}Repository])],
  controllers: [{{MODULE_NAME}}Controller],
  providers: [{{MODULE_NAME}}Service],
})
export class {{MODULE_NAME}}Module {}