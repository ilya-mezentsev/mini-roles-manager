import { ListItemRoute, NavigationProps } from './navigation.types';
import ListItem from '@material-ui/core/ListItem';
import { Link } from 'react-router-dom';
import List from '@material-ui/core/List';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';

export const NavigationList = (props: NavigationProps<ListItemRoute>) => (
    <>
        <List>
            {
                props.routes.map((r, index) => (
                    <ListItem
                        button
                        component={Link}
                        to={r.path}
                        key={`list_item_route_${index}`}
                    >
                        {
                            r.iconComponent &&
                            <ListItemIcon>
                                { <r.iconComponent/> }
                            </ListItemIcon>
                        }
                        <ListItemText primary={r.name} />
                    </ListItem>
                ))
            }
        </List>
    </>
);
