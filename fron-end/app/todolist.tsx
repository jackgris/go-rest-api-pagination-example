'use client';
import { useEffect, useState } from 'react';
import {Todo,TodoProp} from "./todo";

const todosProps: TodoProp[] = [
    { id: 1, name: "Todo 1" },
    { id: 2, name: "Todo 2" },
    { id: 3, name: "Todo 3" },
];

export const TodoList = (): JSX.Element => {
    const [todos, setTodos] = useState<TodoProp[]>([]);
   
    useEffect(() => {
        setTodos(todosProps);
    }, []);
    
    return (
        <>
            <div>List TODOS</div>
            {todos.map((todo) => <Todo key={todo.id} {...todo}/>)}    
        </>
    )
}