import Mock from 'mockjs';
import qs from 'query-string';
import setupMock from '@/utils/setupMock';

setupMock({
  setup: () => {
    Mock.mock(new RegExp('/api/v1/businesses'), () => {
      return [
        {
            title: '评论点赞',
            business_id: 'COMMENT_LIKE',
            type: 'SIMPLE_VOTE',
        },
        {
            title: '评论点踩',
            business_id: 'COMMENT_UNLIKE',
            type: 'SIMPLE_VOTE',
        },
      ];
    });

    Mock.mock(new RegExp('/api/v1/vote_system'), () => {
        return [
          {
              title: '评论点赞',
              business_id: 'COMMENT_LIKE',
              type: 'SIMPLE_VOTE',
          },
          {
              title: '评论点踩',
              business_id: 'COMMENT_UNLIKE',
              type: 'SIMPLE_VOTE',
          },
        ];
      });
  
  },
});
