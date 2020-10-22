import { ApiProperty } from '@nestjs/swagger';
import { IsIn, IsOptional } from 'class-validator';
import {} from '../{{MODULE_NAME_LOWER}}.enum';

// @NeedsManualIntervention
export class Get{{MODULE_NAME}}FilterDto {
    // @IsOptional()
    // @ApiProperty({
    //   description: 'name of department to filter list of members',
    //   required: false,
    // })
    // name: string;
}