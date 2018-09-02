import React from 'react';
import Card from '@material-ui/core/Card';
import CardHeader from '@material-ui/core/CardHeader';
import CardContent from '@material-ui/core/CardContent';
import Typography from '@material-ui/core/Typography';

export default () => (
    <Card>
        <CardHeader title="Error" />
        <CardContent>
            <Typography component="h2">
                404: Page not found
            </Typography>
        </CardContent>
    </Card>
);
