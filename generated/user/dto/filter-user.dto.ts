import { ApiProperty } from '@nestjs/swagger';
import { IsIn, IsOptional } from 'class-validator';
import {} from '../user.enum';

// @NeedsManualIntervention
export class GetUserFilterDto {
    // @IsOptional()
    // @ApiProperty({
    //   description: 'name of department to filter list of members',
    //   required: false,
    // })
    // name: string;
}