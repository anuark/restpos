import React from 'react';
import { List, Create, Edit, Datagrid, TextField, DateField, ImageField, Show, SimpleShowLayout, SimpleForm, ImageInput } from 'react-admin';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardMedia from '@material-ui/core/CardMedia';
import Button from '@material-ui/core/Button';
// import CardContent from '@material-ui/core/CardContent';
import CardHeader from '@material-ui/core/CardHeader';
import PersonIcon from '@material-ui/icons/Person';
import Avatar from '@material-ui/core/Avatar';

const cardStyle = {
    width: 300,
    minHeight: 300,
    margin: '0.5em',
    display: 'inline-block',
    verticalAlign: 'top'
};

const cardMediaStyle = {
    height: '150px'
};

const ItemGrid = ({ ids, data, basePath }) => (
    <div style={{margin: '1em'}}>
        {ids.map(id => 
            <Card key={id} style={cardStyle}>
                {/* <CardMedia
                    image={data[id].image.url}
                    title={data[id].image.desc}
                /> */}
                <CardMedia image={data[id].image.url} style={cardMediaStyle} />
                <CardHeader
                    // title={<TextField record={data[id]} source="name" />}
                    title={data[id].name}
                    subtitle={data[id].created_at}
                    // subtitle={<DateField record={data[id]} source="created_at" />}
                    // avatar={<Avatar icon={<PersonIcon />} />}
                />
                <CardActions>
                    <Button size="small" color="primary">
                    Enable
                    </Button>
                    <Button size="small" color="primary">
                    Edit
                    </Button>
                </CardActions>
            </Card>
        )}
    </div>
);

ItemGrid.defaultProps = {
    data: {},
    ids: []
};

export const ItemList = (props) => (
    <List {...props} title="Items">
        <ItemGrid />
    </List>
);

export const ItemCreate = (props) => (
    <Create {...props}>
        <SimpleForm>
            <TextField source="name" />
            <ImageInput source="image" accept="image/*" />
        </SimpleForm>
    </Create>
);

export const ItemEdit = (props) => (
    <Edit {...props}>
        <ImageInput source="image" accept="image/*">
            <TextField source="name" />
            <ImageField source="image" />
        </ImageInput>
    </Edit>
);

export const ItemShow = (props) => (
    <Show {...props}>
        <SimpleShowLayout>
            <TextField source="id" />
            <TextField source="name" />
            <ImageField source="image.url" title="image.desc" />
        </SimpleShowLayout>
    </Show>
);
