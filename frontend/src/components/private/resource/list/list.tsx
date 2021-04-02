import { ListProps } from './list.types';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import ListItemSecondaryAction from '@material-ui/core/ListItemSecondaryAction';
import { IconButton } from '@material-ui/core';
import { Delete, Edit } from '@material-ui/icons';
import Divider from '@material-ui/core/Divider';

export const ResourcesList = (props: ListProps) => {
    if (props.resources.length > 0) {
        return (
            <List component="nav">
                {
                    props.resources.map(resource => (
                        <>
                            <ListItem>
                                <ListItemText primary={`${resource.id}: ${resource.title || 'No title'}`} />
                                <ListItemSecondaryAction>
                                    <IconButton edge="end">
                                        <Edit onClick={() => props.tryEdit(resource)} />
                                    </IconButton>
                                    <IconButton edge="end">
                                        <Delete onClick={() => props.tryDelete(resource)} />
                                    </IconButton>
                                </ListItemSecondaryAction>
                            </ListItem>
                            <Divider/>
                        </>
                    ))
                }
            </List>
        );
    } else {
        return (
            <h3>No resources created yet</h3>
        );
    }
}
