import React from 'react';
// also exported from '@storybook/react' if you can deal with breaking changes in 6.1
import { Story, Meta } from '@storybook/react/types-6-0';

import {Block, BlockProps} from '../components/Block';

export default {
  title: 'Example/BlockButton',
  component: Block,
  argTypes: {
    backgroundColor: { control: 'color' },
  },
} as Meta;

const Template: Story<BlockProps> = (args) => <Block {...args} />;

export const Large = Template.bind({});
Large.args = {
  children: 'My Block Button',
};

export const Medium = Template.bind({});
Medium.args = {
  children: 'My Block Button',
  size: 'medium'
};

export const Small = Template.bind({});
Small.args = {
  children: 'My Block Button',
  size: 'small'
};

export const Active = Template.bind({});
Active.args = {
  children: 'Active',
  size: 'small',
  active: true
};

export const Passive = Template.bind({});
Passive.args = {
  children: 'Passive',
  size: 'small'
};