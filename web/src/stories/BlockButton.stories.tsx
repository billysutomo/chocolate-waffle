import React from 'react';
// also exported from '@storybook/react' if you can deal with breaking changes in 6.1
import { Story, Meta } from '@storybook/react/types-6-0';

import {BlockButton, BlockButtonProps} from '../components/BlockButton';

export default {
  title: 'Example/BlockButton',
  component: BlockButton,
  // argTypes: {
  //   backgroundColor: { control: 'color' },
  // },
} as Meta;

const Template: Story<BlockButtonProps> = (args) => <BlockButton {...args} />;

export const Large = Template.bind({});
Large.args = {
  label: 'My Block Button',
};

export const Medium = Template.bind({});
Medium.args = {
  label: 'My Block Button',
  size: 'medium'
};

export const Small = Template.bind({});
Small.args = {
  label: 'My Block Button',
  size: 'small'
};