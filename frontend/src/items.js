import React from 'react';
import { List, Create, Edit, DataGrid, TextField, ImageField } from 'react-admin';

export const ItemList = () => (
    <List {...props}>
        <DataGrid>
            <TextField source="id" />
            <TextField source="name" />
            <ImageField source="image_url" />
        </DataGrid>
    </List>
);

export const ItemCreate = () => (
    <Create>

    </Create>
);

export const ItemEdit = () => (
    <Edit>

    </Edit>
);
