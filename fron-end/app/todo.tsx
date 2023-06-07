export interface TodoProp  {
    id: number;
    name: string;
}

export const Todo = ({id, name}:TodoProp): JSX.Element => {
    
    return (
        <div>{name}</div>
    )
};
