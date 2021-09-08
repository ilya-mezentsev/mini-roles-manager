interface EditableField {
    name: string;
    value: string;
    label: string;
    onChange: (newValue: string) => void;
    disabled?: boolean;
}

export interface TextFieldsProps {
    fields: EditableField[];
}
