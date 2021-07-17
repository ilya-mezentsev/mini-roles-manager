import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import ListItemSecondaryAction from '@material-ui/core/ListItemSecondaryAction';
import { IconButton } from '@material-ui/core';
import { Delete, Edit } from '@material-ui/icons';
import Divider from '@material-ui/core/Divider';
import List from '@material-ui/core/List';

import { RolesVersionListProps } from './list.types';

export const RolesVersionList = (props: RolesVersionListProps) => {
    if (props.rolesVersions.length > 0) {
        return (
            <List>
                {
                    props.rolesVersions.map(rv => (
                        <>
                            <ListItem>
                                <ListItemText primary={`${rv.id}: ${rv.title || 'No title'}`} />
                                <ListItemSecondaryAction>
                                    <IconButton edge="end">
                                        <Edit onClick={() => props.tryEdit(rv)} />
                                    </IconButton>
                                    <IconButton edge="end">
                                        <Delete onClick={() => props.tryDelete(rv)} />
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
        // Incorrect situation. But must be handled
        return (
            <h3>No resources versions created yet</h3>
        );
    }
};
