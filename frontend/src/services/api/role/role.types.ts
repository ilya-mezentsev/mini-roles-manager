
export interface Role {
    id: string;
    title?: string;
    permissions?: string[];
    extends?:  string[];
}
