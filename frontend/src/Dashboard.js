import React from 'react';
import Card from '@material-ui/core/Card';
import CardHeader from '@material-ui/core/CardHeader';
import CardContent from '@material-ui/core/CardContent';
import Typography from '@material-ui/core/Typography';

export default () => (
    <Card>
        <CardHeader title="Hola"></CardHeader>
        <CardContent>
            <Typography component="h1">
                World
            </Typography>
        </CardContent>
    </Card>
);
