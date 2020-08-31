FROM scratch

ADD main /
ADD templates /templates

CMD ["/main"]