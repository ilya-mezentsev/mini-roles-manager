import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import ListItemSecondaryAction from '@material-ui/core/ListItemSecondaryAction';
import { IconButton } from '@material-ui/core';
import { Delete, Edit } from '@material-ui/icons';
import Divider from '@material-ui/core/Divider';
import Chip from '@material-ui/core/Chip';

import { ListProps } from './list.types';

export const RolesList = (props: ListProps) => {
    if (props.roles.length > 0) {
        return (
            <List>
                {
                    props.roles.map(role => (
                        <>
                            <ListItem>
                                <ListItemText
                                    primary={`${role.id}: ${role.title || 'No title'}`}
                                    secondary={
                                        role.extends?.map((roleId, index) => (
                                            <Chip
                                                label={roleId}
                                                key={`extends_role_chip_${roleId}_${index}`}
                                                color="primary"
                                            />
                                        ))
                                    }
                                />
                                <ListItemSecondaryAction>
                                    <IconButton edge="end">
                                        <Edit onClick={() => props.tryEdit(role)} />
                                    </IconButton>
                                    <IconButton edge="end">
                                        <Delete onClick={() => props.tryDelete(role)} />
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
            <h3>No roles created yet</h3>
        );
    }
}
