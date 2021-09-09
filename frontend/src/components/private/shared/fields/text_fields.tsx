import { TextField } from '@material-ui/core';

import { TextFieldsProps } from './text_fields.types';


export const TextFields = (props: TextFieldsProps) => (
    <>
        {
            props.fields.map((field, index) => (
                <TextField
                    key={`generic_edit_field_${field.name}_${index}`}
                    margin="dense"
                    label={field.label}
                    fullWidth
                    disabled={field.disabled}
                    value={field.value}
                    onChange={e => field.onChange((e.target as HTMLInputElement).value)}
                />
            ))
        }
    </>
);
